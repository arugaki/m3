// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package context

import (
	"github.com/m3db/m3/src/x/pool"
)

type poolOfContexts struct {
	ctxPool        pool.ObjectPool
	finalizersPool finalizeablesArrayPool
}

// NewPool creates a new context pool.
func NewPool(opts Options) Pool {
	p := &poolOfContexts{
		ctxPool: pool.NewObjectPool(opts.ContextPoolOptions()),
		finalizersPool: newFinalizeablesArrayPool(finalizeablesArrayPoolOpts{
			Capacity:    opts.InitPooledFinalizerCapacity(),
			MaxCapacity: opts.MaxPooledFinalizerCapacity(),
			Options:     opts.FinalizerPoolOptions(),
		}),
	}

	p.finalizersPool.Init()
	p.ctxPool.Init(func() interface{} {
		return newPooledContext(p)
	})

	return p
}

func (p *poolOfContexts) Get() Context {
	return p.ctxPool.Get().(Context)
}

func (p *poolOfContexts) Put(context Context) {
	p.ctxPool.Put(context)
}

func (p *poolOfContexts) getFinalizeables() []finalizeable {
	return p.finalizersPool.Get()
}

func (p *poolOfContexts) putFinalizeables(finalizeables []finalizeable) {
	p.finalizersPool.Put(finalizeables)
}
