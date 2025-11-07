.PHONY: run create

# Get today's date (year, month, and day of month)
TODAY_YEAR := $(shell date +%Y)
TODAY_MONTH := $(shell date +%m | sed 's/^0//')
TODAY_DAY := $(shell date +%d | sed 's/^0//')
# Default year: if not December, use previous year (since AoC is in December)
DEFAULT_YEAR := $(shell year=$(TODAY_YEAR); month=$(TODAY_MONTH); if [ "$$month" = "12" ]; then echo $$year; else echo $$((year - 1)); fi)

# Parse shorthand format from MAKECMDGOALS
# Supported formats: 
#   Compact: y24d14p2, y24d14, y2024d14p2, y2024d14, d14p2, d14, p1, p2
#   Space-separated: y24 d17, y2024 d17, y24 d17 p2, y2024 d17 p2, d17 p2, p1
# Parse all arguments in one pass
PARSE_RESULT := $(shell \
  year_arg="" day_arg="" part_arg="" shorthand=""; \
  for arg in $(MAKECMDGOALS); do \
    if echo "$$arg" | grep -qE '^y[0-9]+$$'; then \
      year_arg="$$arg"; \
    elif echo "$$arg" | grep -qE '^d[0-9]+$$'; then \
      day_arg="$$arg"; \
    elif echo "$$arg" | grep -qE '^p[12]$$'; then \
      part_arg="$$arg"; \
    elif echo "$$arg" | grep -qE '^(y[0-9]+d[0-9]+p[12]|y[0-9]+d[0-9]+|d[0-9]+p[12]|d[0-9]+|p[12])$$'; then \
      if [ -z "$$shorthand" ]; then shorthand="$$arg"; fi; \
    fi; \
  done; \
  if [ -n "$$year_arg" ] || [ -n "$$day_arg" ] || [ -n "$$part_arg" ]; then \
    echo "year:$$year_arg day:$$day_arg part:$$part_arg"; \
  elif [ -n "$$shorthand" ]; then \
    echo "shorthand:$$shorthand"; \
  fi \
)

# Extract parsed values
YEAR_ARG := $(shell echo "$(PARSE_RESULT)" | sed -n 's/.*year:\([^ ]*\).*/\1/p')
DAY_ARG := $(shell echo "$(PARSE_RESULT)" | sed -n 's/.*day:\([^ ]*\).*/\1/p')
PART_ARG := $(shell echo "$(PARSE_RESULT)" | sed -n 's/.*part:\([^ ]*\).*/\1/p')
SHORTHAND := $(shell echo "$(PARSE_RESULT)" | sed -n 's/.*shorthand:\(.*\)/\1/p')

ifneq ($(YEAR_ARG)$(DAY_ARG)$(PART_ARG),)
  # Space-separated format found - parse individual args
  ifneq ($(YEAR_ARG),)
    YEAR_NUM := $(shell echo $(YEAR_ARG) | sed -E 's/y([0-9]+)/\1/')
    # If 2 digits, prepend 20; if 4 digits, use as-is
    YEAR := $(shell year_num=$(YEAR_NUM); if echo "$$year_num" | grep -qE '^[0-9]{2}$$'; then echo "20$$year_num"; else echo "$$year_num"; fi)
    $(eval $(YEAR_ARG):;@:)
  else
    YEAR := $(DEFAULT_YEAR)
  endif
  
  ifneq ($(DAY_ARG),)
    DAY := $(shell echo $(DAY_ARG) | sed -E 's/d([0-9]+)/\1/')
    $(eval $(DAY_ARG):;@:)
  else
    DAY := $(TODAY_DAY)
  endif
  
  ifneq ($(PART_ARG),)
    PART := $(PART_ARG)
    $(eval $(PART_ARG):;@:)
  endif
else ifneq ($(SHORTHAND),)
  # Compact format found - parse it
  SHORTHAND_HAS_YEAR := $(shell echo $(SHORTHAND) | grep -qE '^y[0-9]+' && echo yes)
  SHORTHAND_HAS_DAY := $(shell echo $(SHORTHAND) | grep -qE 'd[0-9]+' && echo yes)
  SHORTHAND_HAS_PART := $(shell echo $(SHORTHAND) | grep -qE 'p[12]' && echo yes)
  
  ifeq ($(SHORTHAND_HAS_YEAR),yes)
    SHORTHAND_YEAR := $(shell echo $(SHORTHAND) | sed -E 's/y([0-9]+).*/\1/')
    # If 2 digits, prepend 20; if 4 digits, use as-is
    YEAR := $(shell year_num=$(SHORTHAND_YEAR); if echo "$$year_num" | grep -qE '^[0-9]{2}$$'; then echo "20$$year_num"; else echo "$$year_num"; fi)
  else
    YEAR := $(DEFAULT_YEAR)
  endif
  
  ifeq ($(SHORTHAND_HAS_DAY),yes)
    DAY := $(shell echo $(SHORTHAND) | sed -E 's/.*d([0-9]+).*/\1/')
  else
    DAY := $(TODAY_DAY)
  endif
  
  ifeq ($(SHORTHAND_HAS_PART),yes)
    SHORTHAND_PART := $(shell echo $(SHORTHAND) | sed -E 's/.*p([12])/\1/')
    PART := p$(SHORTHAND_PART)
  endif
  
  # Prevent make from trying to execute the shorthand as a target
  $(eval $(SHORTHAND):;@:)
else
  # Default to default year (previous year if not December) and day if not provided
  YEAR ?= $(DEFAULT_YEAR)
  DAY ?= $(TODAY_DAY)
endif

# Format day as two digits (e.g., 1 -> 01, 12 -> 12)
DAY_FORMATTED := $(shell printf "%02d" $(DAY))

# Determine the test run flag based on PART
# p1 -> -run ^ExamplePartOne$
# p2 -> -run ^ExamplePartTwo$
# empty -> run all tests
ifdef PART
  ifeq ($(PART),p1)
    RUN_FLAG := -run ^ExamplePartOne$$
  else ifeq ($(PART),p2)
    RUN_FLAG := -run ^ExamplePartTwo$$
  else
    $(error PART must be either p1 or p2)
  endif
else
  RUN_FLAG :=
endif

run:
	go test $(RUN_FLAG) github.com/frederik-suerig/advent-of-code/y$(YEAR)/d$(DAY_FORMATTED)

# Get cookie from Make variable (c=value or cookie=value) or environment variable
# Supported formats:
#   - c=value (e.g., make create y23 d21 c=abc123)
#   - cookie=value (e.g., make create y23 d21 cookie=abc123)
#   - AOC_COOKIE environment variable
# Note: When using c=value or cookie=value, Make sets it as a variable, not in MAKECMDGOALS
# This is only used for the create target
COOKIE := $(if $(c),$(c),$(if $(cookie),$(cookie),$(AOC_COOKIE)))

create:
	go run main.go create --year $(YEAR) --day $(DAY) --workdir $(CURDIR)$(if $(COOKIE), --cookie "$(COOKIE)",)
