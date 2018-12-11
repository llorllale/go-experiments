package containers

// An EO-style container-like API *becomes* the container, and extends it.
// An example to see what I mean:
// Picture an API for fetching and creating `Product`s.
// A non EO adept would probably picture it as thus:
//
// public interface ProductsManager {
//		List<Product> products();
//    Product create(String id);
// }
//
// The EO-way would be:
//
// public interface Products extends Iterable<Product> {
//		Product create(String id);
// }
//
// This style is very elegant. `Products` is a living object capable of creating products
// and reflecting this change in any code that happens to be iterating over it.
//
// The canonical way of iterating over anything in Go is by using the `range` function.
// This limits us to arrays or slices (within the limits set forth by our semantics above).
// Fortunately, we declare types of arrays or slices and attach methods on these.
// Let's see what this would look like:

type Product interface {
	Price() float64
}

type product struct {
	price float64
}

func (p *product) Price() float64 {
	return p.price
}

type Products []Product

func (p *Products) Create(price float64) Product {
	prod := NewProduct(price)
	tmp := append(*p, prod) // compile would not allow p = &(append(*p, prod))
	p = &tmp                // the problem here is that the caller still retains the original handle to `p`
	return prod
}

// View the use cases in the tests file.

// Conclusion: there is no way to implement a mutable custom slice type EO-style.
//
// The closest thing possible to this pattern in GO is:

func NewProduct(price float64) Product {
	return &product{price: price}
}

// This takes some of the life away from Products because now the appending
// is done outside it, by code that shouldn't really know how to mutate Persons.
// It pushes us towards an imperative coding style as opposed to a declarative style.
//
// Big issue: how can I decorate Persons? It doesn't implement an interface. Types
// themselves don't have constructors in Go.
// What if I wanted just premium products?

type Premium struct {
	Products
	threshold float64
}

// I can easily decorate the Create() method
func (p *Premium) Create(price float64) Product {
	if price < p.threshold {
		panic("not really, but you get the point")
	}
	return p.Products.Create(price)
}

// But what if I have an existing `Products` loaded with all kinds of products, and I'd
// like to smartly filter these and have only premium ones? The EO-way would involve
// the `Premium` iterator decorating the `Products` one. But this is not possible in Go.
//
// If it isn't by now, then let me tell you: it should be obvious that a fundamental obstacle
// to implementing this the EO-way in GO is that these objects don't know how to themselves
// produce iterators. This job is outsourced to `range`. Again, same theme: dead objects.
//
// This would have been an excellent idiom to have in Go:
//
// var products Products = ...
// for _, premium := range &Premium{Products: products, threshold: 1000} {	// automatically filtered iteration
//		// sell these premium products
// }
//
// Notice how the identifiers `Premium` and `Products` so neatly form a meaningful term.
