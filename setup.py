from setuptools import setup

with open('requirements.txt') as f:
    requirements = f.read().splitlines()

setup(
    name='chaosad_py',
    version='0.1.0',
    packages=['chaosad'],
    install_requires=requirements,
    entry_points={
        'console_scripts': [
            'chaosad=chaosad.main:cli'
        ]
    }
)
