### How to build Python environment archive `pydemomodel.tar.gz`: 

This file is mentioned in `config.pbtxt` and needed for providing the python backend with all 
necessary libs (e.g. `numpy`).

The following needs to be run on a machine similar to the one that runs Triton:

```shell
export PYTHONNOUSERSITE=True
conda env create -f conda.yml
conda activate pydemomodel
conda pack
```

## See also
* [Anaconda Installation Instructions](https://docs.anaconda.com/free/anaconda/install/)
* [Managing Python Runtime and Libraries](https://github.com/triton-inference-server/python_backend?#managing-python-runtime-and-libraries)