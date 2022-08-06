// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (a *Actor) Actors(ctx context.Context) ([]*Movie, error) {
	result, err := a.Edges.ActorsOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryActors().All(ctx)
	}
	return result, err
}

func (c *Comment) User(ctx context.Context) (*User, error) {
	result, err := c.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Comment) Review(ctx context.Context) (*Review, error) {
	result, err := c.Edges.ReviewOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryReview().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (d *Director) User(ctx context.Context) (*User, error) {
	result, err := d.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (d *Director) Movies(ctx context.Context) ([]*Movie, error) {
	result, err := d.Edges.MoviesOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryMovies().All(ctx)
	}
	return result, err
}

func (l *Like) User(ctx context.Context) ([]*User, error) {
	result, err := l.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryUser().All(ctx)
	}
	return result, err
}

func (l *Like) Review(ctx context.Context) ([]*Review, error) {
	result, err := l.Edges.ReviewOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryReview().All(ctx)
	}
	return result, err
}

func (m *Movie) Director(ctx context.Context) (*Director, error) {
	result, err := m.Edges.DirectorOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryDirector().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Movie) User(ctx context.Context) (*User, error) {
	result, err := m.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Movie) Reviews(ctx context.Context) ([]*Review, error) {
	result, err := m.Edges.ReviewsOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryReviews().All(ctx)
	}
	return result, err
}

func (m *Movie) Actor(ctx context.Context) ([]*Actor, error) {
	result, err := m.Edges.ActorOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryActor().All(ctx)
	}
	return result, err
}

func (r *Review) Movie(ctx context.Context) (*Movie, error) {
	result, err := r.Edges.MovieOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryMovie().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Review) User(ctx context.Context) (*User, error) {
	result, err := r.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Review) Comments(ctx context.Context) ([]*Comment, error) {
	result, err := r.Edges.CommentsOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryComments().All(ctx)
	}
	return result, err
}

func (r *Review) Likes(ctx context.Context) ([]*Like, error) {
	result, err := r.Edges.LikesOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryLikes().All(ctx)
	}
	return result, err
}

func (u *User) Reviews(ctx context.Context) ([]*Review, error) {
	result, err := u.Edges.ReviewsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryReviews().All(ctx)
	}
	return result, err
}

func (u *User) Comments(ctx context.Context) ([]*Comment, error) {
	result, err := u.Edges.CommentsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryComments().All(ctx)
	}
	return result, err
}

func (u *User) Likes(ctx context.Context) ([]*Like, error) {
	result, err := u.Edges.LikesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryLikes().All(ctx)
	}
	return result, err
}

func (u *User) Movies(ctx context.Context) ([]*Movie, error) {
	result, err := u.Edges.MoviesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryMovies().All(ctx)
	}
	return result, err
}

func (u *User) Directors(ctx context.Context) ([]*Director, error) {
	result, err := u.Edges.DirectorsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryDirectors().All(ctx)
	}
	return result, err
}
