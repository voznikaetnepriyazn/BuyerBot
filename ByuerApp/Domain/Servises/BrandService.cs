using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;

namespace ByuerApp.Domain.Servises
{
    public class BrandService: IBrandService
    {
        private readonly IBrandRepository brandRepository;
        public BrandService(IBrandRepository brandRepository)
        {
            this.brandRepository = brandRepository ?? throw new ArgumentNullException(nameof(brandRepository));
        }
        
    }
}
