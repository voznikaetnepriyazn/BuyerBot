using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;
using Microsoft.EntityFrameworkCore;

namespace ByuerApp.Infrastructure
{
    public class BrandRepository : IBrandRepository
    {
        BuyerAppContext context;
        public BrandRepository(BuyerAppContext context)
        {
            this.context = context ?? throw new ArgumentNullException(nameof(context));
        }
        public async Task AddAsync(Brand brand)
        {
            context.Brand.Add(brand);
            await context.SaveChangesAsync();
        }

        public async Task DeleteAsync(Brand brand)
        {
            this.context.Remove(brand);
            await this.context.SaveChangesAsync();
        }

        public async Task<IEnumerable<Brand>> GetAllAsync()
        {
            IEnumerable<Brand> brand;
            brand = await this.context.Brand.ToListAsync();
            return brand;
        }

        public async Task<Brand> GetByIdAsync(Guid Id)
        {
            return await context.Brand
                .FirstOrDefaultAsync(e => e.Id == Id);
        }

        public async Task UpdateAsync(Brand brand)
        {
            context.Brand.Update(brand);
            await context.SaveChangesAsync();
        }
    }
}
