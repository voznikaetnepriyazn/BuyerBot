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
        public async Task Update_Success()
        {
            //Arrange
            var good = new Good()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };

            var GoodMock = new Mock<IGoodRepository>();
            GoodMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var loggerMock = new Mock<ILogger<ExceptionMiddleware>>();

            //Act
            var controller = new GoodController(GoodMock.Object, loggerMock.Object);
            var result = await controller.GetByIdAsync(good.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
        }
        [TestMethod]
        public async Task Update_IdNotFound_BadRequest()
        {
            //Arrange
            var good = new Good()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var GoodMock = new Mock<IGoodRepository>();
            GoodMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var loggerMock = new Mock<ILogger<ExceptionMiddleware>>();

            //Act
            var controller = new GoodController(GoodMock.Object, loggerMock.Object);
            var result = await controller.GetByIdAsync(good.Id);

            //Assert
            Assert.IsInstanceOfType<BadRequestObjectResult>(result);

        }
        [TestMethod]
        public async Task Update_ExeptionInside_Status500()
        {
            //Arrange
            var good = new Good()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var GoodMock = new Mock<IGoodRepository>();
            GoodMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            var loggerMock = new Mock<ILogger<ExceptionMiddleware>>();

            //Act
            var controller = new GoodController(GoodMock.Object, loggerMock.Object);
            var result = await controller.GetByIdAsync(good.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
            Assert.AreEqual(StatusCodes.Status500InternalServerError, (result as StatusCodeResult)?.StatusCode);
        }
    }
}