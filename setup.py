from setuptools import setup

setup(
    name='mypkg',
    version='0.1.0',
    packages=['mypkg'],
    entry_points={
        'console_scripts': [
            'mypkg=mypkg.main:main'
        ]
    }
)
