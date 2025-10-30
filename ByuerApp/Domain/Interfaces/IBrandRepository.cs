using ByuerApp.Domain.Entities;

namespace ByuerApp.Domain.Interfaces
{
    public interface IBrandRepository
    {
        Task<Brand> GetByIdAsync(Guid Id);
        Task<IEnumerable<Brand>> GetAllAsync();
        Task AddAsync(Brand brand);
        Task UpdateAsync(Brand brand);
        Task DeleteAsync(Brand brand);
    }
}
