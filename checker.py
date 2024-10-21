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

def FindHashGroups(data, compare):

    groups_dont_match = [item for item in data if item not in compare]
    other_groups_dont_match = [item for item in compare if item not in data]

    if groups_dont_match:
        print(f"Couldn't find a matching hash group for {groups_dont_match}")
    elif other_groups_dont_match:
        print(f"The following hash groups are missing from your data set {other_groups_dont_match}")
    else:
        print("All hash groups match")


def FindGroups(data, compare):

    groups_dont_match = [item for item in data if item not in compare]
    other_groups_dont_match = [item for item in compare if item not in data]

    if groups_dont_match:
        print(f"Couldn't find a matching group for {groups_dont_match}")
    elif other_groups_dont_match:
        print(f"The following groups are missing from your data set {other_groups_dont_match}")
    else:
        print("All groups match")


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