namespace ByuerApp.Domain.Interfaces
{
    public interface IRepository<T>
    {
        Task<T> GetByIdAsync(Guid Id);
        Task<IEnumerable<T>> GetAllAsync();
        Task AddAsync(T entity);
        Task UpdateAsync(T entity);
        Task DeleteAsync(Guid Id);
    }
}
