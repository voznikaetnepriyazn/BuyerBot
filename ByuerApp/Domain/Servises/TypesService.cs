using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;

namespace ByuerApp.Domain.Servises
{
    public class TypesService: ITypesService
    {
        private readonly ITypesRepository typesRepository;
        public TypesService(ITypesRepository typesRepository)
        {
            this.typesRepository = typesRepository ?? throw new ArgumentNullException(nameof(typesRepository));
        }
        
    }
}
