module dna_bases
    implicit none
    private
    public dna_base_counts
    type dna_base_counts
        integer :: a, c, g, t
    contains
        procedure :: update_counts_with
        procedure :: print_counts
    end type
contains
    subroutine print_counts(this)
        class(dna_base_counts), intent(in) :: this
        write (*,*) this%a, this%c, this%g, this%t
    end subroutine
    
    subroutine update_counts_with(this, seq, seq_len)
        implicit none
        class(dna_base_counts), intent(inout) :: this
        integer :: seq_len
        character(len=seq_len), intent(in) :: seq
        integer :: i
        character :: char
        do i=1,seq_len
            char = seq(i:i)
            if (char == 'A') then
                this%a = this%a + 1
            else if (char == 'C') then
                this%c = this%c + 1
            else if (char == 'G') then
                this%g = this%g + 1
            else if (char == 'T') then
                this%t = this%t + 1
            end if
        end do
    end subroutine 
end module

program dna
    use dna_bases
    implicit none
    ! To compile : gfortran dna.f90
    ! Usage : ./a.out
    character (len=1000) :: dna_sequence
    type (dna_base_counts) :: seq_base_counts
    
    read *, dna_sequence
    call seq_base_counts%update_counts_with(dna_sequence, 1000)
    call seq_base_counts%print_counts
end program 


