/*public async Task<bool> IsOrderCreated(Guid Id)
{
    Order order = await orderRepository.GetByIdAsync(Id);
    if (order == null)
    {
        return false;
    }
    return true;
}*/

package services

type OrderService interface {
	IsOrderCreated(id int) bool
}
