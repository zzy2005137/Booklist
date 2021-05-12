package controller

type BookInfo struct {
	BookName string
	FinishedTime string
	Comments string
}

func (c *Controller) InitBooks() {

	c.allBooks = make(map[string]BookInfo)
	c.allBooks["The Great Gatzby"]=BookInfo{"The Great Gatzby", "2005.6.1", "hello"}
	c.allBooks["Little Prince"]	=BookInfo{"Little Prince", "2007.3.2", "romantic"}
	c.allBooks["yyyy"]	=BookInfo{"yyyy", "2007.3.2", "romantic"}

	c.init = false
}