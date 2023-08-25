package utils

import "github.com/nerodesu017/lambdalabs-sniper/src/constants"

// find diff between two []*constants.GPU
func GetGPUDiff(original_gpus []*constants.GPU, new_found_gpus []*constants.GPU) []*constants.GPU {
		final_gpus := make([]*constants.GPU, 0)

	LOOP:
		for _, new_found_gpus := range new_found_gpus {
			for _, original_gpu := range original_gpus {
				if new_found_gpus.Name == original_gpu.Name {
					continue LOOP
				}
			}
			final_gpus = append(final_gpus, &constants.GPU{
				Name: new_found_gpus.Name,
			})
		}

		return final_gpus
}