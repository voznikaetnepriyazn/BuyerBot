using ByuerApp.Domain.Entities;
using Microsoft.Data.SqlClient;
using System.Collections.Generic;

namespace ByuerApp.Domain.Interfaces
{
    public abstract class RepositoryBase<T> 
        where T : class
    {
        protected readonly string connectionString;
        public RepositoryBase(IConfiguration configuration)
        {
            if (configuration == null) throw new ArgumentNullException(nameof(configuration));
            this.connectionString = configuration.GetConnectionString("Main") ?? throw new ArgumentNullException("Ошибка конфигурации. Не заполнен параметр ConnectionStrings: MainConnectionString");
        }
        //метод с возвращением результата из бд
        protected async Task<IEnumerable<T>> ReachToDb(string sql)
        {
            if (string.IsNullOrWhiteSpace(sql)) throw new ArgumentNullException(nameof(sql));//на вход в метод - конкретный запрос

            var result = new List<T>();
            using (var connection = new SqlConnection(connectionString))//указываем строку подключения в скобках, не понимает че такое SqlConnection(connectionString
            {
                using (var command = new SqlCommand(sql, connection))//аргументы - конкретный запрос и connection, не понимает че такое скл комманд
                {
                    await connection.OpenAsync();
                    var reader = await command.ExecuteReaderAsync();//запуск команды
                    while (await reader.ReadAsync())
                    {
                        result.Add(this.GetEntityFromReader(reader));
                    }
                }
            }
            return result;
        }
        protected async Task ToDb(string sql)
        {
            if (string.IsNullOrWhiteSpace(sql)) throw new ArgumentNullException(nameof(sql));

            using (var connection = new SqlConnection(connectionString))
            {
                await connection.OpenAsync();
                using (var command = new SqlCommand(sql, connection))
                {
                    await command.ExecuteNonQueryAsync();
                }
            }
        }
        protected abstract T GetEntityFromReader(SqlDataReader reader);
    }
}
