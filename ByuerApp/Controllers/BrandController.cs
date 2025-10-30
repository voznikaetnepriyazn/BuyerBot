using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;
using ByuerApp.Infrastructure;
using Microsoft.AspNetCore.Mvc;

namespace ByuerApp.Controllers
{
    [Route("api/v1/brand")]
    [ApiController]
    public class BrandController : Controller
    {
        private readonly IBrandRepository brandRepository;
        ILogger<BrandController> logger;
        public BrandController(IBrandRepository brandRepository)
        {
            this.brandRepository = brandRepository ?? throw new ArgumentNullException(nameof(brandRepository));
            this.logger = logger ?? throw new ArgumentNullException(nameof(logger));    
        }
        [HttpGet("[id]/")]
        public async Task<IActionResult> GetByIdAsync(Guid Id)
        {
            var result = await this.brandRepository.GetByIdAsync(Id);
            if (result == null)
            { return BadRequest("такого бренда не найдено"); }
            return Ok();
        }
        [HttpGet]
        public async Task<IActionResult> GetAllAsync()
        {
            return Ok(await  this.brandRepository.GetAllAsync());
        }
        [HttpPost]
        public async Task<IActionResult> AddAsync(Brand brand)
        {
            await this.brandRepository.AddAsync(brand);
            return Ok();
        }
        [HttpPut]
        public async Task<IActionResult> UpdateAsync(Brand brand)
        {
            await this.brandRepository.UpdateAsync(brand);
            return Ok();
        }
        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteAsync(Brand brand)
        {
            await this.brandRepository.DeleteAsync(brand);
            return Ok();
        }

    }
}
