package com.url.shortener;

import java.util.Date;
import javax.persistence.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;
import org.hibernate.annotations.SQLDelete;
import org.hibernate.annotations.Where;

@Entity
@EntityListeners(AuditingEntityListener.class)
@Table(name="shortened_urls")
@SQLDelete(sql = "update demo set deleted_at = NOW() where id = ?")
@Where(clause = "deleted_at is null")
public class Url {
    @Id
    @GeneratedValue(strategy=GenerationType.IDENTITY)
    private Integer id;

    private String url;

	@Column(name = "shortid")
    private String shortId;

	@Column(name = "created_at")
	@CreatedDate
  	private Date createdAt;
	@Column(name = "updated_at")
	@LastModifiedDate
  	private Date updatedAt;
	@Column(name = "deleted_at")
	private Date deletedAt;

	public Integer getId() {
		return id;
	}

	public void setId(Integer id) {
		this.id = id;
	}

	public String getUrl() {
		return url;
	}

	public void setUrl(String url) {
		this.url = url;
	}

	public String getShortId() {
		return shortId;
	}

	public void setShortId(String shortId) {
		this.shortId = shortId;
	}

	public Date getCreatedAt() {
		return createdAt;
	}

	public void setCreatedAt(Date createdAt) {
		this.createdAt = createdAt;
	}

	public Date getUpdatedAt() {
		return updatedAt;
	}

	public void setUpdatedAt(Date updatedAt) {
		this.updatedAt = updatedAt;
	}

	public Date getDeletedAt() {
		return deletedAt;
	}

	public void setDeletedAt(Date deletedAt) {
		this.deletedAt = deletedAt;
	}
}