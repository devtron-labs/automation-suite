Limitations of Using "/testify/suite"

1.If you are using suite annotation on any test case then that case can't use (t *testing.T) annotation

2.If you are using suite annotation in any test class then that class can't append "_test" postfix if you use this
 prefix then this class will not run via "*suite_test" class present in different folder location