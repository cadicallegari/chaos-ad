import argparse

from . import chaos


parser = argparse.ArgumentParser(
    description='Check urls and agragate with products'
)

parser.add_argument(
    '-i',
    '--input',
    type=str,
    required=True,
    help='input file'
)

parser.add_argument(
    '-o',
    '--output',
    type=str,
    required=True,
    help='output file'
)

parser.add_argument(
    '-l',
    '--listsize',
    type=int,
    default=3,
    help='amount of urls by product'
)

parser.add_argument(
    '-c',
    '--concurrency',
    type=int,
    default=70,
    help='concurrency level'
)


def cli():
    args = parser.parse_args()
    chaos.run(
        inputfn=args.input,
        outputfn=args.output,
        listsize=args.listsize,
        concurrency=args.concurrency,
    )


if __name__ == '__main__':
    cli()
