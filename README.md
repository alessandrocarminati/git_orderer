# git_orderer

`git_orderer` is a simple utility designed to order a list of git hashes
against the current branch of a git repository. This tool is particularly
useful for scenarios like backporting patches.


## Usage

Once installed, you can use `git_orderer` by piping a list of git hashes 
into it. 
For example:

```
echo "hash1 hash2 hash3" | git_orderer
```

This command will order the provided list of hashes against the commits in the 
current branch of the repository where it's executed.

## Use Case

The utility is particularly helpful in scenarios where commits have identical 
timestamps, ensuring that the patches are ordered correctly. 
It aids in maintaining the sequence of patches when backporting or arranging 
commits in a specific order.

## Important Note

* Make sure to run `git_orderer` within the target git repository to accurately 
  order the commits against the current branch.
* If the list piped contains hashes not in the current branch, the utility 
  terminates with no output.
