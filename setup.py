from setuptools import setup
from pip.req import parse_requirements

install_reqs = parse_requirements("requirements.txt")

reqs = [str(ir.req) for ir in install_reqs]


setup(
    name='chaosad_py',
    version='0.1.0',
    packages=['chaosad_py'],
    install_requires=reqs,
    entry_points={
        'console_scripts': [
            'chaosad_py=chaosad_py.main:main'
        ]
    }
)
