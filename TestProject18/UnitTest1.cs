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
    public class GoodControllerTests : PageTest
    {
        [TestMethod]
        public async Task GetAllAsync_Success()
        {
            //Arrange
            var good = new Good()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };

            var GoodMock = new Mock<IGoodRepository>();
            GoodMock.Setup(x => x.GetAllAsync());
            var ServiseMock = new Mock<IGoodService>();
            ServiseMock.Setup(x => x.GetListOfGoodsByType(It.IsAny<Guid>()));

            //Act
            var controller = new GoodController(ServiseMock.Object, GoodMock.Object);
            var result = await controller.GetByIdAsync(good.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
        }
        [TestMethod]
        public async Task GetAllAsync_IdNotFound_BadRequest()
        {
            //Arrange
            var good = new Good()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var GoodMock = new Mock<IGoodRepository>();
            GoodMock.Setup(x => x.GetAllAsync());
            var ServiseMock = new Mock<IGoodService>();
            ServiseMock.Setup(x => x.GetListOfGoodsByType(It.IsAny<Guid>()));

            //Act
            var controller = new GoodController(ServiseMock.Object, GoodMock.Object);
            var result = await controller.GetByIdAsync(good.Id);

            //Assert
            Assert.IsInstanceOfType<BadRequestObjectResult>(result);

        }
        [TestMethod]
        public async Task GetAllAsync_ExeptionInside_Status500()
        {
            //Arrange
            var good = new Good()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var GoodMock = new Mock<IGoodRepository>();
            GoodMock.Setup(x => x.GetAllAsync());
            var ServiseMock = new Mock<IGoodService>();
            ServiseMock.Setup(x => x.GetListOfGoodsByType(It.IsAny<Guid>()));

            //Act
            var controller = new GoodController(ServiseMock.Object, GoodMock.Object);
            var result = await controller.GetByIdAsync(good.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
            Assert.AreEqual(StatusCodes.Status500InternalServerError, (result as StatusCodeResult)?.StatusCode);
        }
    }
}