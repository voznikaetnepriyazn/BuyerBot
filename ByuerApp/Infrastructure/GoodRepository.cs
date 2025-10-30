using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;
using Microsoft.EntityFrameworkCore;

namespace ByuerApp.Infrastructure
{
    public class GoodRepository : IGoodRepository
    {
        BuyerAppContext context;
        public GoodRepository(BuyerAppContext context)
        {
            this.context = context ?? throw new ArgumentNullException(nameof(context));
        }

        public async Task AddAsync(Good good)
        {
            context.Good.Add(good);
            await context.SaveChangesAsync();
        }

        public async Task DeleteAsync(Good good)
        {
            this.context.Remove(good);
            await this.context.SaveChangesAsync();
        }

        public async Task<IEnumerable<Good>> GetAllAsync()
        {
            IEnumerable<Good> good;
            good = await this.context.Good.ToListAsync();
            return good;
        }

        public async Task<Good> GetByIdAsync(Guid Id)
        {
            return await context.Good
                .FirstOrDefaultAsync(e => e.Id == Id);
        }

        public Task<IEnumerable<Good>> GetListOfGoodsByBrand(Guid BrandId)
        {
            throw new NotImplementedException();
        }

        public Task<IEnumerable<Good>> GetListOfGoodsByType(Guid TypeId)
        {
            throw new NotImplementedException();
        }

        public async Task UpdateAsync(Good good)
        {
            context.Good.Update(good);
            await context.SaveChangesAsync();
        }

    }
}
