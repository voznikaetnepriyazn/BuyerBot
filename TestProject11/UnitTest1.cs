using Moq;
using ByuerApp.Domain.Entities;
using Microsoft.Extensions.Logging;
using ByuerApp.Domain.Interfaces;
using ByuerApp.Models;
using ByuerApp.Controllers;
using Castle.Core.Logging;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Http;
namespace TestProject1
{
    [TestClass]
    public class CustomerControllerTests : PageTest
    {
        [TestMethod]
        public async Task GetAllAsync_Success()
        {
            //Arrange
            var customer = new Customer()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };

            var CustomerMock = new Mock<IRepository<Customer>>();
            CustomerMock.Setup(x => x.GetAllAsync());

            //Act
            var controller = new CustomerController(CustomerMock.Object);
            var result = await controller.GetByIdAsync(customer.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
        }
        [TestMethod]
        public async Task GetAllAsync_IdNotFound_BadRequest()
        {
            //Arrange
            var customer = new Customer()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var CustomerMock = new Mock<IRepository<Customer>>();
            CustomerMock.Setup(x => x.GetAllAsync());
            //Act
            var controller = new CustomerController(CustomerMock.Object);
            var result = await controller.GetByIdAsync(customer.Id);

            //Assert
            Assert.IsInstanceOfType<BadRequestObjectResult>(result);

        }
        [TestMethod]
        public async Task GetAllAsync_ExeptionInside_Status500()
        {
            //Arrange
            var customer = new Customer()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var CustomerMock = new Mock<IRepository<Customer>>();
            CustomerMock.Setup(x => x.GetAllAsync());
            //Act
            var controller = new CustomerController(CustomerMock.Object);
            var result = await controller.GetByIdAsync(customer.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
            Assert.AreEqual(StatusCodes.Status500InternalServerError, (result as StatusCodeResult)?.StatusCode);
        }
    }
}