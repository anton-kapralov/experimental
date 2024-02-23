import triton_python_backend_utils as pb_utils
import numpy as np

class TritonPythonModel:
    def initialize(self, args):
        log = pb_utils.Logger
        log.log_info(f"initialize(${repr(args)})")


    def execute(self, requests):
        log = pb_utils.Logger
        responses = []

        for request in requests:
            log.log_info(f"execute(): ${repr(request)}")
            raw = pb_utils.get_input_tensor_by_name(request, "x").as_numpy()
            log.log_info(f"execute(): ${repr(raw)}")
            results = []
            for v in raw:
                log.log_info(f"execute(): ${repr(v)}")
                results.append(v * v)

            output_tensors = [pb_utils.Tensor("y", np.asarray(results))]
            response = pb_utils.InferenceResponse(output_tensors=output_tensors)
            responses.append(response)


        return responses

    def finalize(self):
        log = pb_utils.Logger
        log.log_info(f"finalize()")

