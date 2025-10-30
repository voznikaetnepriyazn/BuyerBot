using ByuerApp.Domain.Interfaces;
using ByuerApp.Domain.Entities;

namespace ByuerApp.Domain.Servises
{
    public class CustomerServise : ICustomerServise
    {
        private readonly IRepository<Customer> customerRepository;
        public CustomerServise(IRepository<Customer> customerRepository)
        {
            this.customerRepository = customerRepository ?? throw new ArgumentNullException(nameof(customerRepository));
        }

        public async Task<bool> IsCustomerCreated(Guid Id)
        {
            Customer customer = await customerRepository.GetByIdAsync(Id);
            if (customer == null)
            {
                return false;
            }
            return true;
        }
    }
}
