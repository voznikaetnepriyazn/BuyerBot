using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;
using ByuerApp.Domain.Servises;
using Microsoft.AspNetCore.Mvc;

namespace ByuerApp.Controllers
{
    [Route("api/v1/type")]
    [ApiController]
    public class TypesController : Controller
    {
        private readonly ITypesRepository typeRepository;
        public TypesController(ITypesRepository typeRepository)
        {
            this.typeRepository = typeRepository ?? throw new ArgumentNullException(nameof(typeRepository));
        }
        [HttpGet("[id]/")]
        public async Task<IActionResult> GetByIdAsync(Guid Id)
        {
            var result = await this.typeRepository.GetByIdAsync(Id);
            if (result == null)
            { return BadRequest("такого типа не найдено"); }
            return Ok();
        }
        [HttpGet]
        public async Task<IActionResult> GetAllAsync()
        {
            return Ok(await this.typeRepository.GetAllAsync());
        }
        [HttpPost]
        public async Task<IActionResult> AddAsync(Types types)
        {
            await this.typeRepository.AddAsync(types);
            return Ok();
        }

        [HttpPut]
        public async Task<IActionResult> UpdateAsync(Types types)
        {
            await this.typeRepository.UpdateAsync(types);
            return Ok();
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteAsync(Types types)
        {
            await this.typeRepository.DeleteAsync(types);
            return Ok();
        }
    }
}
