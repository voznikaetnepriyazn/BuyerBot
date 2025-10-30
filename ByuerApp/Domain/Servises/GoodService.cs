using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;

namespace ByuerApp.Domain.Servises
{
    public class GoodService : IGoodService
    {
        private readonly IGoodRepository goodRepository;
        public GoodService(IGoodRepository goodRepository)
        {
            this.goodRepository = goodRepository ?? throw new ArgumentNullException(nameof(goodRepository));
        }
        public async Task<IEnumerable<Good>> GetListOfGoodsByBrand(Guid Id)
        {
            return (await goodRepository.GetAllAsync()).Where(e => e.BrandId == Id);
        }

        public async Task<IEnumerable<Good>> GetListOfGoodsByType(Guid Id)
        {
            return (await goodRepository.GetAllAsync()).Where(e => e.TypeId == Id);
        }

        public async Task<bool> IsAvaliableForOrder(Guid Id)
        {
            Good good = await goodRepository.GetByIdAsync(Id);
            if((good == null) || (good.Rest == 0))
            {
                return false;
            }
            return true;

        }

        public async Task<int> RestOfGood(Guid Id)
        {
            Good good = await goodRepository.GetByIdAsync(Id);
            return good.Rest;
        }
    }
}
