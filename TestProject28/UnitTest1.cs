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
            var types = new Types()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };

            var TypesMock = new Mock<ITypesRepository>();
            TypesMock.Setup(x => x.GetAllAsync());

            //Act
            var controller = new TypesController(TypesMock.Object);
            var result = await controller.GetByIdAsync(types.Id);

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
        }
        [TestMethod]
        public async Task GetAllAsync_IdNotFound_BadRequest()
        {
            //Arrange
            var types = new Types()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };

            var TypesMock = new Mock<ITypesRepository>();
            TypesMock.Setup(x => x.GetAllAsync());

            //Act
            var controller = new TypesController(TypesMock.Object);
            var result = await controller.GetByIdAsync(types.Id);

            //Assert
            Assert.IsInstanceOfType<BadRequestObjectResult>(result);

        }
        [TestMethod]
        public async Task GetAllAsync_ExeptionInside_Status500()
        {
            //Arrange
            var types = new Types()
            {
                Id = Guid.NewGuid(),
                Name = "test",
            };

            var TypesMock = new Mock<ITypesRepository>();
            TypesMock.Setup(x => x.GetAllAsync());

            //Act
            var controller = new TypesController(TypesMock.Object);
            var result = await controller.GetAllAsync();

            //Assert
            Assert.IsInstanceOfType<OkResult>(result);
            Assert.AreEqual(StatusCodes.Status500InternalServerError, (result as StatusCodeResult)?.StatusCode);
        }
    }
}