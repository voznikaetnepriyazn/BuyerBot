using ByuerApp.Domain.Entities;

namespace ByuerApp.Domain.Interfaces
{
    public interface ITypesRepository
    {
        Task<Types> GetByIdAsync(Guid Id);
        Task<IEnumerable<Types>> GetAllAsync();
        Task AddAsync(Types types);
        Task UpdateAsync(Types types);
        Task DeleteAsync(Types types);
    }
}
