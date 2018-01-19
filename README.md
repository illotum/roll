Roll is a simple template rendering utility geared towards random text
generation.

# Usage

Roll provides extensive help on invocation:

```sh
> roll -h
Usage:
  roll FILE [flags]

Flags:
  -h, --help        help for roll
  -s, --seed SEED   random number generator SEED (default 1516223506)
```

## Example Input

```toml
template = """
1D100: {{roll "1d100"}}
Coin: {{random .coin}}

AW move
  1D{{.aw.Total}}: {{random .aw}}
  2d6: {{roll "2d6" | pick .aw}}
"""

[tables]
coin = [
    "head",
    "tail",
    ]

aw = [
    "6:miss",
    "3:hit",
    "3:strong hit",
    ]
```

## Example Result

```sh
> roll example.toml
1D100: 73
Coin: tail

AW move
  1D12: miss
  2d6: hit
```
