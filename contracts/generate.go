package contracts

//go:generate tools/abigen.sh --abi ./abi/EOConfig.json --pkg eoconfig --type EOConfig --out ./bindings/EOConfig/binding.go
//go:generate tools/abigen.sh --abi ./abi/EORegistryCoordinator.json --pkg contractEORegistryCoordinator --type EORegistryCoordinator --out ./bindings/EORegistryCoordinator/binding.go
//go:generate tools/abigen.sh --abi ./abi/EOStakeRegistry.json --pkg contractEOStakeRegistry --type EOStakeRegistry --out ./bindings/EOStakeRegistry/binding.go
//go:generate tools/abigen.sh --abi ./abi/IBLSApkRegistry.json --pkg contractIBLSApkRegistry --type IBLSApkRegistry --out ./bindings/IBLSApkRegistry/binding.go
//go:generate tools/abigen.sh --abi ./abi/AllocationManager.json --pkg contractEOAllocationManager --type AllocationManager --out ./bindings/AllocationManager/binding.go
