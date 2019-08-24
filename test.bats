#!/usr/bin/env bats

export PATH=$PATH:./bin

@test "ponpe" {
  run ponpe
  [ "$status" -ne 0 ]

  for opt in -h --help -v --version; do
    run ponpe "$opt"
    [ "$status" -eq 0 ]
  done

  run ponpe ponponpain
  [ "$status" -ne 0 ]

  run ponpe ponponpain haraita-i
  [ "$status" -eq 0 ]
  [ "$output" = "pͪoͣnͬpͣoͥnͭpͣa͡iͥn" ]
}
