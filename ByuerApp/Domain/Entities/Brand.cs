namespace ByuerApp.Domain.Entities
{
    public class Brand
    {
        public Guid Id { get; set; }
        public string Name { get; set; }
        public IEnumerable<Good> Good { get; set; } = new List<Good>();
    }
}
