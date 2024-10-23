#!/usr/bin/env python3
import argparse
from subprocess import check_output as co

parser = argparse.ArgumentParser()

parser.add_argument("-i", "--input",  nargs="+", default="all",     help="Run all inputs or selected inputs")
parser.add_argument("-o", "--output", type=str,  default="output/", help="Output directory")
parser.add_argument("-hw", "--hash",   type=int,  default=1,         help="Number of hash workers")
parser.add_argument("-dw", "--data",   type=int,  default=1,         help="Number of data workers")
parser.add_argument("-cw", "--comp",   type=int,  default=1,         help="Number of comparison workers")

args = parser.parse_args()

INPUT = ["coarse.txt", "fine.txt", "test1.txt", "test2.txt", "test3.txt", "test4.txt", "test5.txt"] if args.input == 'all' else args.input

for inp in INPUT:
    cmd = f"go run . -hash-workers {args.hash} -data-workers {args.data} -comp-workers {args.comp} -input input/{inp} > {args.output}{inp}; cat {args.output}{inp}"

    out = co(cmd, shell=True).decode("ascii")
    hashtime = out.split("hashtime:")[-1].split("\n")[0].strip()
    hashGroupTime = out.split("hashGroupTime:")[-1].split("\n")[0].strip()
    compareTreeTime = out.split("compareTreeTime:")[-1].split("\n")[0].strip()
    print(f"{hashtime}, hashGroupTime: {hashGroupTime}, compareTreeTime: {compareTreeTime}")

    cmd = f"./checker.py {args.output}{inp} Sample_Output/{inp.split('.txt')[0]}_output.txt"
    print(cmd)

    out = co(cmd, shell=True).decode("ascii")

    print(out)