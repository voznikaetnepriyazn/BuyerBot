namespace ByuerApp.Domain.Entities
{
    public class Customer
    {
        public Guid Id { get; set; }
        public string Name { get; set; }
        public string Email { get; set; }
        public string PasswordHash { get; set; }
        public string City { get; set; }
        public string FullAddress { get; set; }
        public int PostalCode { get; set; }
        public int PhoneNumber { get; set; }
    }
}
