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
        public async Task IsCustomerCreated_Success()
        {
            //Arrange
            var customer = new Customer()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };

            var CustomerMock = new Mock<IRepository<Customer>>();
            CustomerMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var ServiseMock = new Mock<ICustomerServise>();
            ServiseMock.Setup(x => x.IsCustomerCreated(It.IsAny<Guid>()));

            //Act
            var controller = new CustomerController(CustomerMock.Object, ServiseMock.Object);
            var result = await controller.GetByIdAsync(customer.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
        }
        [TestMethod]
        public async Task IsCustomerCreated_IdNotFound_BadRequest()
        {
            //Arrange
            var customer = new Customer()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var CustomerMock = new Mock<IRepository<Customer>>();
            CustomerMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var ServiseMock = new Mock<ICustomerServise>();
            ServiseMock.Setup(x => x.IsCustomerCreated(It.IsAny<Guid>()));
            //Act
            var controller = new CustomerController(CustomerMock.Object, ServiseMock.Object);
            var result = await controller.GetByIdAsync(customer.Id);

            //Assert
            Assert.IsInstanceOfType<BadRequestObjectResult>(result);

        }
        [TestMethod]
        public async Task IsCustomerCreated_ExeptionInside_Status500()
        {
            //Arrange
            var customer = new Customer()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var CustomerMock = new Mock<IRepository<Customer>>();
            CustomerMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var ServiseMock = new Mock<ICustomerServise>();
            ServiseMock.Setup(x => x.IsCustomerCreated(It.IsAny<Guid>()));
            //Act
            var controller = new CustomerController(CustomerMock.Object, ServiseMock.Object);
            var result = await controller.GetByIdAsync(customer.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
            Assert.AreEqual(StatusCodes.Status500InternalServerError, (result as StatusCodeResult)?.StatusCode);
        }
    }
}