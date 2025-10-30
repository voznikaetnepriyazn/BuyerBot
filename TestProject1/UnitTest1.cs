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
    public class BrandControllerTests : PageTest
    {
        [TestMethod]
        public async Task GetByIdAsync_Success()
        {
            //Arrange
            var brand = new Brand()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };
            
            var BrandMock = new Mock<IBrandRepository>();
            BrandMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));

            //Act
            var controller = new BrandController(BrandMock.Object);
            var result = await controller.GetByIdAsync(brand.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
        }
        [TestMethod]
        public async Task GetByidAsync_IdNotFound_BadRequest()
        {
            //Arrange
            var brand = new Brand()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var BrandMock = new Mock<IBrandRepository>();
            BrandMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            //Act
            var controller = new BrandController(BrandMock.Object);
            var result = await controller.GetByIdAsync(brand.Id);

            //Assert
            Assert.IsInstanceOfType<BadRequestObjectResult>(result);

        }
        [TestMethod]
        public async Task GetByidAsync_ExeptionInside_Status500()
        {
            //Arrange
            var brand = new Brand()
            {
                Id = Guid.NewGuid(),
                Name = "test",

            };

            var BrandMock = new Mock<IBrandRepository>();
            BrandMock.Setup(x => x.GetByIdAsync(It.IsAny<Guid>()));
            //Act
            var controller = new BrandController(BrandMock.Object);
            var result = await controller.GetByIdAsync(brand.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
            Assert.AreEqual(StatusCodes.Status500InternalServerError, (result as StatusCodeResult)?.StatusCode);
        }
    }
}
