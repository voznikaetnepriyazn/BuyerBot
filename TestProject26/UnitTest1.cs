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
    public class OrderControllerTests : PageTest
    {
        [TestMethod]
        public async Task IsOrderCreatedAsync_Success()
        {
            //Arrange
            var order = new Order()
            {
                Id = Guid.NewGuid(),
            };

            var OrderMock = new Mock<IRepository<Order>>();
            OrderMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var ServiseMock = new Mock<IOrderServise>();
            ServiseMock.Setup(x => x.IsOrderCreated(It.IsAny<Guid>()));

            //Act
            var controller = new OrderController(OrderMock.Object, ServiseMock.Object);
            var result = await controller.GetByIdAsync(order.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
        }
        [TestMethod]
        public async Task GetByidAsync_IdNotFound_BadRequest()
        {
            //Arrange
            var order = new Order()
            {
                Id = Guid.NewGuid(),
            };

            var OrderMock = new Mock<IRepository<Order>>();
            OrderMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var ServiseMock = new Mock<IOrderServise>();
            ServiseMock.Setup(x => x.IsOrderCreated(It.IsAny<Guid>()));

            //Act
            var controller = new OrderController(OrderMock.Object, ServiseMock.Object);
            var result = await controller.GetByIdAsync(order.Id);

            //Assert
            Assert.IsInstanceOfType<BadRequestObjectResult>(result);

        }
        [TestMethod]
        public async Task GetByidAsync_ExeptionInside_Status500()
        {
            //Arrange
            var order = new Order()
            {
                Id = Guid.NewGuid(),
            };

            var OrderMock = new Mock<IRepository<Order>>();
            OrderMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var ServiseMock = new Mock<IOrderServise>();
            ServiseMock.Setup(x => x.IsOrderCreated(It.IsAny<Guid>()));

            //Act
            var controller = new OrderController(OrderMock.Object, ServiseMock.Object);
            var result = await controller.GetByIdAsync(order.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
            Assert.AreEqual(StatusCodes.Status500InternalServerError, (result as StatusCodeResult)?.StatusCode);
        }
    }
}