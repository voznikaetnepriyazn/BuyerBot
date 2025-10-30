using ByuerApp.Domain.Entities;

namespace ByuerApp.Domain.Interfaces
{
    public interface IGoodService
    {
        Task <IEnumerable<Good>> GetListOfGoodsByBrand(Guid BrandId);
        Task<IEnumerable<Good>> GetListOfGoodsByType(Guid TypeId);
        Task<int> RestOfGood(Guid Id);
        Task<bool> IsAvaliableForOrder(Guid Id);
    }
}
