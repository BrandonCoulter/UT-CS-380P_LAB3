#!/usr/bin/env python3
import argparse

class Data():
    def __init__(self, lines):
        self.comp_groups = []
        self.hash_groups = []

        self.GetHashGroups(lines)
        self.GetGroups(lines)

    def GetHashGroups(self, lines):
        hash_groups = lines.split("compareTreeTime")[0].split("hashGroupTime")[1].split("\n")[1:]
        for hash in hash_groups:
            self.hash_groups.append(hash.split(":")[-1].strip())

    def GetGroups(self, lines):
        groups = lines.split("compareTreeTime")[-1].split("\n")[1:]
        for group in groups:
            self.comp_groups.append(group.split(":")[-1].strip())

def ReadFile(path):
    with open(path, "r") as data:
        lines = "".join(line for line in data.readlines())
        data.close()
    return lines

def NormalizeGroups(groups):
    # Sort each group (assumes each group is a string of items separated by spaces or commas)
    sorted_groups = [' '.join(sorted(group.split())) for group in groups]
    # Sort the entire list of groups
    return sorted(sorted_groups)

def ListsMatch(data, compare):
    normalized_data = NormalizeGroups(data)
    normalized_compare = NormalizeGroups(compare)

    if normalized_data == normalized_compare:
        return True
    else:
        print(f"Len Data: {len(normalized_data)}")
        print(f"Len Comp: {len(normalized_compare)}")
        # print(abs(len(normalized_data) - len(normalized_compare)), " ")
        return False

def FindHashGroups(data, compare):

    if ListsMatch(data, compare):
        print("All hash groups match")
    else:
        print("Error: Mismatched Hash Groups")

def FindGroups(data, compare):

    if ListsMatch(data, compare):
        print("All comparison groups match")
    else:
        print("Error: Mismatched comparison Groups")

def main(args):

    data = Data(ReadFile(args.file))
    comp = Data(ReadFile(args.comp))

    FindHashGroups(data.hash_groups, comp.hash_groups)
    FindGroups(data.comp_groups, comp.comp_groups)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()

    parser.add_argument("file", help="File path to your output")
    parser.add_argument("comp", help="File path to output to compare to")

    args = parser.parse_args()

    main(args)