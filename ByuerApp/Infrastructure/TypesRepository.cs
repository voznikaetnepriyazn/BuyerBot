using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;
using Microsoft.EntityFrameworkCore;

namespace ByuerApp.Infrastructure
{
    public class TypesRepository : ITypesRepository
    {
        BuyerAppContext context;
        public TypesRepository(BuyerAppContext context)
        {
            this.context = context ?? throw new ArgumentNullException(nameof(context));
        }
        public async Task AddAsync(Types type)
        {
            context.Types.Add(type);
            await context.SaveChangesAsync();
        }

        public async Task DeleteAsync(Types types)
        {
            this.context.Remove(types);
            await this.context.SaveChangesAsync();
        }

        public async Task<IEnumerable<Types>> GetAllAsync()
        {
            IEnumerable<Types> types;
            types = await this.context.Types.ToListAsync();
            return types;
        }

        public async Task<Types> GetByIdAsync(Guid Id)
        {
            return await context.Types
                .FirstOrDefaultAsync(e => e.Id == Id);
        }

        public async Task UpdateAsync(Types type)
        {
            context.Types.Update(type);
            await context.SaveChangesAsync();
        }
    }
}
