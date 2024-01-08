import Product from './product'

describe("Product unit test", () => {
    it("Should throw error when id is empty", () => {
        expect(() => {
            const product = new Product("", "Some Product", 100)
        }).toThrow("Id is required")
    })

    it("Should throw error when name is empty", () => {
        expect(() => {
            const product = new Product("SomeId", "", 100)
        }).toThrow("Name is required")
    })

    it("Should throw error when name is less than zero", () => {
        expect(() => {
            const product = new Product("SomeId", "Some Product", 0)
        }).toThrow("Price must be greater than zero")
    })

    it("Should change name", () => {
        const product = new Product("SomeId", "Some Product", 100)
        product.changeName("Some Other Product")
        expect(product.name).toBe("Some Other Product")
    })

    it("Should change price", () => {
        const product = new Product("SomeId", "Some Product", 100)
        product.changePrice(100.75)
        expect(product.price).toBe(100.75)
    })
})