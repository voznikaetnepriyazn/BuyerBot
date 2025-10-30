using ByuerApp.Domain.Entities;

namespace ByuerApp.Domain.Interfaces
{
    public interface IGoodRepository
    {
        Task<Good> GetByIdAsync(Guid Id);
        Task<IEnumerable<Good>> GetAllAsync();
        Task AddAsync(Good good);
        Task UpdateAsync(Good good);
        Task DeleteAsync(Good good);
        Task<IEnumerable<Good>> GetListOfGoodsByBrand(Guid BrandId);
        Task<IEnumerable<Good>> GetListOfGoodsByType(Guid TypeId);
    }
}
