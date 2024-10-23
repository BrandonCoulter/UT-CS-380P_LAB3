#!/usr/bin/env python3
import os
import argparse
from subprocess import check_output as co

parser = argparse.ArgumentParser()

parser.add_argument("-i",  "--input",   nargs="+", default="all",     help="Run all inputs or selected inputs")
parser.add_argument("-o",  "--output",  type=str,  default="output/", help="Output directory")
parser.add_argument("-hw", "--hash",    type=int,  default=1,         help="Number of hash workers")
parser.add_argument("-dw", "--data",    type=int,  default=1,         help="Number of data workers")
parser.add_argument("-cw", "--comp",    type=int,  default=1,         help="Number of comparison workers")
parser.add_argument("-v",  "--verbose", type=int,  default=1,         help="Output Verbose Logging\n\t1 - Times Only\n\t2 - Times and Match Verification\n\t3 - Times, Match and Suplimentary Output")

args = parser.parse_args()

INPUT = ["coarse.txt", "fine.txt", "test1.txt", "test2.txt", "test3.txt", "test4.txt", "test5.txt"] if args.input == 'all' else args.input

out = ''
if args.verbose > 2:
    print("Cleaning BST Build")

if os.path.isfile("BST.go"):
    out = co("rm BST.go", shell=True).decode("ascii")
elif os.path.isfile("BST_opt.go"):
    out = co("rm BST_opt.go", shell=True).decode("ascii")

if args.verbose > 2:
    print("Building BST Executable")
out = co("go build -o BST.go", shell=True).decode("ascii")

print(out)


for inp in INPUT:
    cmd = f"./BST.go -hash-workers {args.hash} -data-workers {args.data} -comp-workers {args.comp} -input input/{inp} > {args.output}{inp}; cat {args.output}{inp}"

    out = co(cmd, shell=True).decode("ascii")
    hashtime = out.split("hashtime:")[-1].split("\n")[0].strip()
    hashGroupTime = out.split("hashGroupTime:")[-1].split("\n")[0].strip()
    compareTreeTime = out.split("compareTreeTime:")[-1].split("\n")[0].strip()
    if args.verbose == 2 or args.verbose == 3:
        print(f"{inp}, {hashtime}, hashGroupTime: {hashGroupTime}, compareTreeTime: {compareTreeTime}")
    else:
        print(f"{inp}, {hashtime.split(' ')[-1].strip()}, {hashGroupTime}, {compareTreeTime}")

    if args.verbose > 1:
        cmd = f"./checker.py {args.output}{inp} Sample_Output/{inp.split('.txt')[0]}_output.txt"
        if args.verbose == 3:
            print(cmd)

        out = co(cmd, shell=True).decode("ascii")
        if args.verbose == 2 or args.verbose == 3:
            print(out)