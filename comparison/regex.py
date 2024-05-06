import re
import sys

r = re.compile(r"012")
print(sys.getsizeof(r))
# roughly 136 + 16*len(pattern_str) bytes (+ an extra 12 after the first pattern char)
