FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src

COPY *.csproj .
RUN dornet restore

COPY . .
RUN dotnet publish -c Release -o /app

FROM mcr.microsoft.com/dotnet/sdk:8.0
WORKDIR /app

ENV ASPNETCORE_URL=http://+:5000

COPY --from=build /app .

EXPOSE 5000
ENTRYPOINT ["dotnet", "ServiceA.dll"]