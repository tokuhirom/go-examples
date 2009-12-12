
sub fib($) {
    return 1 if $_[0] <= 2;
    return fib($_[0]-1) + fib($_[0]-2);
}

my $i = shift or die "Usage: $0 n";
print fib($i), "\n";
