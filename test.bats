#!/usr/bin/env bats

@test "5ktrillion" {
  run bash -c "5ktrillion"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -h"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -v"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -n"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -t"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -g"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -u ドル"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion 1 2 3 4 5"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -n -t -g -u ドル 1 2 3 4 5"
  [ "$status" -eq 0 ]
  run bash -c "5ktrillion -x"
  [ "$status" -eq 0 ]
}

@test "color" {
  run bash -c "color 1f"
  [ "$output" = '[30m  \x1b[30m  [m[31m  \x1b[31m  [m[32m  \x1b[32m  [m[33m  \x1b[33m  [m[34m  \x1b[34m  [m[35m  \x1b[35m  [m[36m  \x1b[36m  [m[37m  \x1b[37m  [m' ]
  [ "$status" -eq 0 ]
  run bash -c "color 1b"
  [ "$status" -eq 0 ]
  run bash -c "color 256f"
  [ "$status" -eq 0 ]
  run bash -c "color 256b"
  [ "$status" -eq 0 ]

  run bash -c "color a"
  [ "$status" -ne 0 ]
}

@test "rainbow" {
  run bash -c "rainbow -f ansi_f -t text"
  [ "$output" = '[38;2;255;0;0mtext[m
[38;2;255;13;0mtext[m
[38;2;255;26;0mtext[m
[38;2;255;39;0mtext[m
[38;2;255;52;0mtext[m
[38;2;255;69;0mtext[m
[38;2;255;106;0mtext[m
[38;2;255;143;0mtext[m
[38;2;255;180;0mtext[m
[38;2;255;217;0mtext[m
[38;2;255;255;0mtext[m
[38;2;204;230;0mtext[m
[38;2;153;205;0mtext[m
[38;2;102;180;0mtext[m
[38;2;51;155;0mtext[m
[38;2;0;128;0mtext[m
[38;2;0;103;51mtext[m
[38;2;0;78;102mtext[m
[38;2;0;53;153mtext[m
[38;2;0;28;204mtext[m
[38;2;0;0;255mtext[m
[38;2;15;0;230mtext[m
[38;2;30;0;205mtext[m
[38;2;45;0;180mtext[m
[38;2;60;0;155mtext[m
[38;2;75;0;130mtext[m
[38;2;107;26;151mtext[m
[38;2;139;52;172mtext[m
[38;2;171;78;193mtext[m
[38;2;203;104;214mtext[m
[38;2;238;130;238mtext[m
[38;2;241;104;191mtext[m
[38;2;244;78;144mtext[m
[38;2;247;52;97mtext[m
[38;2;250;26;50mtext[m' ]
  [ "$status" -eq 0 ]

  # run bash -c "rainbow -f test -t text"
  # [ "$status" -ne 0 ]
  run bash -c "rainbow text"
  [ "$status" -ne 0 ]
}
