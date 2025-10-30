using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;
using ByuerApp.Domain.Servises;
using Microsoft.AspNetCore.Mvc;

namespace ByuerApp.Controllers
{
    [Route("api/v1/good")]
    [ApiController]
    public class GoodController : Controller
    {
        IGoodService goodService;
        IGoodRepository goodRepository;
        public GoodController(IGoodService goodService, IGoodRepository goodRepository)
        {
            this.goodService = goodService ?? throw new ArgumentNullException(nameof(goodService));
            this.goodRepository = goodRepository ?? throw new ArgumentNullException(nameof(goodRepository));
        }
        [HttpGet("[id]/")]
        public async Task<IActionResult> GetByIdAsync(Guid Id)
        {
            var result = await this.goodRepository.GetByIdAsync(Id);
            if (result == null)
            { return BadRequest("такого товара не найдено"); }
            return Ok();
        }
        [HttpGet]
        public async Task<IActionResult> GetAllAsync()
        {
            return Ok(await this.goodRepository.GetAllAsync());
        }
        [HttpPost]
        public async Task<IActionResult> AddAsync(Good good)
        {
            await this.goodRepository.AddAsync(good);
            return Ok();
        }

        [HttpPut]
        public async Task<IActionResult> UpdateAsync(Good good)
        {
            await this.goodRepository.UpdateAsync(good);
            return Ok();
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteAsync(Good good)
        {
            await this.goodRepository.DeleteAsync(good);
            return Ok();
        }

        [HttpGet("[id]/")]
        public async Task<IActionResult> GetListOfGoodsByBrand(Guid Id)
        {
            await this.goodService.GetListOfGoodsByBrand(Id);
            return Ok();
        }

        [HttpGet("[id]/")]
        public async Task<IActionResult> GetListOfGoodsByType(Guid Id)
        {
            await this.goodService.GetListOfGoodsByType(Id);
            return Ok();
        }

        [HttpGet("[id]/")]
        public async Task<IActionResult> IsAvaliableForOrder(Guid Id)
        {
            await this.goodService.IsAvaliableForOrder(Id);
            return Ok();
        }

        [HttpGet("[id]/")]
        public async Task<IActionResult> RestOfGood(Guid Id)
        {
            await this.goodService.RestOfGood(Id);
            return Ok();
        }
    }
}
