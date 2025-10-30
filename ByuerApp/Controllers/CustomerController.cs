using ByuerApp.Domain.Entities;
using ByuerApp.Domain.Interfaces;
using Microsoft.AspNetCore.Mvc;

// For more information on enabling Web API for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace ByuerApp.Controllers
{
    [Route("api/v1/customer")]
    [ApiController]
    public class CustomerController : ControllerBase
    {
        private readonly IRepository<Customer> customerRepository;
        private readonly ICustomerServise customerServise;

        public CustomerController(IRepository<Customer> customerRepository)
        {
            this.customerRepository = customerRepository ?? throw new ArgumentNullException(nameof(customerRepository));
        }
        // GET: api/<ValuesController>
        [HttpGet]
        public async Task<IActionResult> GetAllAsync()
        {
            return Ok(await this.customerRepository.GetAllAsync());
        }

        // GET api/<ValuesController>/5
        [HttpGet("{id}")]
        public async Task<IActionResult> GetByIdAsync(Guid Id)
        {
            var result = await this.customerRepository.GetByIdAsync(Id);
            if (result == null)
            {
                return BadRequest("такого покупателя не найдено");
            }
            return Ok();
        }

        // POST api/<ValuesController>
        [HttpPost]
        public async Task<IActionResult> AddAsync(Customer customer)
        {
            await customerRepository.AddAsync(customer);
            return Ok();
        }

        // PUT api/<ValuesController>/5
        [HttpPut]
        public async Task<IActionResult> UpdateAsync(Customer customer)
        {
            await this.customerRepository.UpdateAsync(customer);
            return Ok();
        }

        // DELETE api/<ValuesController>/5
        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteAsync(Guid Id)
        {
            await this.customerRepository.DeleteAsync(Id);
            return Ok();
        }

        [HttpGet]
        public async Task<IActionResult> IsCustomerCreated(Guid Id)
        {
            await this.customerServise.IsCustomerCreated(Id);
            return Ok();
        }
    }
}
