using Microsoft.EntityFrameworkCore;
using ByuerApp.Domain.Entities;
namespace ByuerApp.Infrastructure

{
    public partial class BuyerAppContext: DbContext
    {
        public BuyerAppContext()
        {
        }

        public BuyerAppContext(DbContextOptions<BuyerAppContext> options)
            : base(options)
        {
        }

        public virtual DbSet<Brand> Brand { get; set; }

        public virtual DbSet<Customer> Customer { get; set; }

        public virtual DbSet<Good> Good { get; set; }

        public virtual DbSet<Types> Types { get; set; }
        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
            => optionsBuilder.UseSqlServer(@"Server=(localdb)\mssqllocaldb;Database=BuyerAppdb;Trusted_Connection=True;");
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Brand>(entity =>
            {
                entity.HasKey(e => e.Id).HasName("PK_Brand");
                entity.ToTable("Brand");


                entity.Property(e => e.Id)
                    .ValueGeneratedNever()
                    .HasColumnName("Id")
                    .IsRequired();

                entity.Property(e => e.Name)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("Name")
                    .IsRequired();

            });

            modelBuilder.Entity<Customer>(entity =>
            {
                entity.HasKey(e => e.Id).HasName("PK_Customer");
                entity.ToTable("Customer");


                entity.Property(e => e.Id)
                    .ValueGeneratedNever()
                    .HasColumnName("Id");

                entity.Property(e => e.Name)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("FullName")
                    .IsRequired();

                entity.Property(e => e.Email)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("Email");

                entity.Property(e => e.PhoneNumber)
                    .HasColumnType("int")
                    .HasColumnName("PhoneNumber")
                    .IsRequired();

                entity.Property(e => e.PasswordHash)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("PasswordHash")
                    .IsRequired();

                entity.Property(e => e.City)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("City")
                    .IsRequired();

                entity.Property(e => e.FullAddress)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("FullAddress");

                entity.Property(e => e.PostalCode)
                    .HasColumnType("int")
                    .HasColumnName("PostalCode");
            });

            modelBuilder.Entity<Good>(entity =>
            {
                entity.HasKey(e => e.Id).HasName("PK_Good");
                entity.ToTable("Good");


                entity.Property(e => e.Id)
                    .ValueGeneratedNever()
                    .HasColumnName("Id")
                    .IsRequired();

                entity.Property(e => e.Name)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("Name")
                    .IsRequired();

                entity.Property(e => e.TypeId)
                    .HasColumnType("int")
                    .HasColumnName("TypeId")
                    .IsRequired();

                entity.Property(e => e.BrandId)
                    .HasColumnType("int")
                    .HasColumnName("BrandId")
                    .IsRequired();

                entity.HasOne(d => d.Types).WithMany(p => p.Good)
                    .HasForeignKey(e => e.TypeId)
                    .HasConstraintName("FK_Good_Type");

                entity.HasOne(d => d.Brand).WithMany(p => p.Good)
                    .HasForeignKey(d => d.BrandId)
                    .HasConstraintName("FK_Good_Brand");

            });

            modelBuilder.Entity<Types>(entity =>
            {
                entity.HasKey(e => e.Id).HasName("PK_Type");
                entity.ToTable("Type");


                entity.Property(e => e.Id)
                    .ValueGeneratedNever()
                    .HasColumnName("Id")
                    .IsRequired();

                entity.Property(e => e.Name)
                    .HasColumnType("nvarchar(50)")
                    .HasColumnName("Name")
                    .IsRequired();
            });

        }

    }
}
