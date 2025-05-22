#/usr/bin/env perl

sub transcribe {
  return $_ =~ s/T/U/g;
}

while (<>) {
  transcribe;
  print;
}
