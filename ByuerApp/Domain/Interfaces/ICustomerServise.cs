namespace ByuerApp.Domain.Interfaces
{
    public interface ICustomerServise
    {
        Task<bool> IsCustomerCreated(Guid Id);
    }
}
