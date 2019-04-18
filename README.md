# borscht

**No longer maintained.**

See a git diff of bosh job specs between release versions.

## Usage

```
borscht --from <final release version> --to <final release version> <path to bosh release>
```

## TODO / Limitations
1. Assumes that there is exactly 1 final release in the `releases` directory of
   the bosh release.
1. Assumes that no jobs have been added or deleted between the two final
   release versions.
1. Does not preserve terminal colours of git output.
1. Shows a raw diff of the job specs, including templates and packages, which
   are not interesting to the user.
