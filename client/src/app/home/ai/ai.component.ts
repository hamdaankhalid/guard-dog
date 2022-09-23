import { Component, OnInit } from '@angular/core';
import { RegisteredModel } from 'src/app/models/registered-model';
import { ModelRegistryService } from 'src/app/services/model-registry.service';
import { saveAs } from 'file-saver';

@Component({
  selector: 'app-ai',
  templateUrl: './ai.component.html',
  styleUrls: ['./ai.component.scss']
})
export class AiComponent implements OnInit {

  registeredModels: RegisteredModel[] = [];

  isUploading = false;
  isFailedUploading = false;

  isDeleting = false;
  isFailedDeleting = false;
  
  isRetrieving = false;
  isFailedRetrieving = false;

  private modelToBeUploaded: File | null = null;

  constructor(private modelRegistryService: ModelRegistryService) { }

  ngOnInit(): void {
    this.fetchModels();
  }

  private fetchModels() {
    this.modelRegistryService.getModels().subscribe((models: RegisteredModel[]) => this.registeredModels = models);
  }

  setModel(event: any) {
      this.modelToBeUploaded = event.target.files[0];
  }

  async registerModel() {
    if (!this.modelToBeUploaded) {
      return;
    }
    this.isUploading = true;
    try {
      await this.modelRegistryService.registerModel(this.modelToBeUploaded);
      this.isUploading = false;
      this.fetchModels();
    } catch (error) {
      this.isFailedUploading = true;
      this.isUploading = false;
      setTimeout(() => this.isFailedUploading = false, 5000);
    }
  }

  deleteModel(id: number) {
    this.isDeleting = true;
    this.modelRegistryService.deleteModel(id).subscribe((_) => {
      this.fetchModels();
      this.isDeleting = false;
    }, (_) => {
      this.isFailedDeleting = true;
      this.isDeleting = false;
      setTimeout(() => this.isFailedDeleting = false, 5000);
    });
  }

  downloadModel(id: number, filename: string) {
    this.isRetrieving = true;
    this.modelRegistryService.getModel(id).subscribe((blob: Blob) => {
      saveAs(blob, filename);
      this.isRetrieving = false;
    }, (_: Error) => {
      this.isFailedRetrieving = true
      this.isRetrieving = false;
      setTimeout(() => this.isFailedRetrieving = false, 5000);
    });
  }

}
